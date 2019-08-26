package microserver

import (
	"context"
	"io/ioutil"
	"os"

	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go"
	"github.com/sirupsen/logrus"
	"google.golang.org/api/option"
)

const defaultFirebaseConfigName = "firebase.json"

func (s *Service) initAuth() error {
	if !mustInitilizeFirebase(s.options) {
		return NoFirebaseOptionsError()
	}

	credentials := s.getFirebaseCredentialsFromOptions()
	s.firebase = initAuthApp(credentials)

	return nil
}

func mustInitilizeFirebase(options *Options) bool {
	return options.db != nil
}

func (s *Service) getFirebaseCredentialsFromOptions() (credentials []byte) {
	firebaseOptions := s.options.firebase
	if firebaseOptions.bucket != "" {
		return s.getCredentialssFromBucket()
	}

	return s.getCredentialsFromFile(firebaseOptions.configPath)
}

func (s *Service) getCredentialssFromBucket() []byte {
	ctx := context.Background()
	bucket := s.options.firebase.bucket
	name := getFirebaseConfigFileNameFromOptions(s.options.firebase)

	storageClient, err := storage.NewClient(ctx)
	if err != nil {
		GetLogger().WithError(err).Fatalf("Can't get storage conection")
	}
	credentials, err := storageClient.Bucket(bucket).Object(name).NewReader(ctx)
	if err != nil {
		GetLogger().WithError(err).Fatal("Can't stablish reader connection")
	}

	buffer, err := ioutil.ReadAll(credentials)
	if err != nil {
		logrus.Fatalf("Can't read credentials: %v", err)
	}

	return buffer
}

func getFirebaseConfigFileNameFromOptions(options *FirebaseOptions) string {
	name := options.name
	if options.name == "" {
		name = defaultFirebaseConfigName
	}

	return name
}

func (s *Service) getCredentialsFromFile(filePath string) []byte {
	file, err := os.Open(filePath)
	if err != nil {
		GetLogger().WithError(err).Fatalf("Can't open credentials file")
	}

	fileStats, err := file.Stat()
	if err != nil {
		GetLogger().WithError(err).Fatalf("Can't stat credentials file")
	}

	buffer := make([]byte, fileStats.Size())
	_, err = file.Read(buffer)

	if err != nil {
		GetLogger().WithError(err).Fatalf("Can't read credentials file")
	}
	return buffer
}

// initAuthApp starts a firebase app with provided credentials
func initAuthApp(firebaseCredentials []byte) *firebase.App {
	context := context.Background()
	opt := option.WithCredentialsJSON(firebaseCredentials)
	firebaseApp, err := firebase.NewApp(context, nil, opt)
	if err != nil {
		GetLogger().WithError(err).Fatal("error initializing app:")
	}
	return firebaseApp
}
