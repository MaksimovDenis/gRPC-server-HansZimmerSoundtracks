package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func YandexStorage() (inmage, interstellar, batman, dune, inception, piratesOfTheCaribbean []string) {

	// Создаем кастомный обработчик эндпоинтов, который для сервиса S3 и региона ru-central1 выдаст корректный URL
	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		if service == s3.ServiceID && region == "ru-central1" {
			return aws.Endpoint{
				PartitionID:   "yc",
				URL:           "https://storage.yandexcloud.net",
				SigningRegion: "ru-central1",
			}, nil
		}
		return aws.Endpoint{}, fmt.Errorf("unknown endpoint requested")
	})

	// Подгружаем конфигрурацию из ~/.aws/*
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithEndpointResolverWithOptions(customResolver))
	if err != nil {
		log.Fatal(err)
	}

	// Создаем клиента для доступа к хранилищу S3
	client := s3.NewFromConfig(cfg)

	// Запрашиваем список бакетов
	result, err := client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
	if err != nil {
		log.Fatal(err)
	}

	for _, bucket := range result.Buckets {
		log.Printf("bucket=%s creation time=%s", aws.ToString(bucket.Name), bucket.CreationDate.Format("2006-01-02 15:04:05 Monday"))
	}
	//Background images storage
	backgoundimages, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String("petprojecthanzzimmer"),
		Prefix: aws.String("backgoundimages/"),
	})
	//Interstellar sountracks storage
	interstellarSoundtrack, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String("petprojecthanzzimmer"),
		Prefix: aws.String("interstellarSoundtrack/"),
	})
	//Batman sountracks storage
	batmanSoundtrack, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String("petprojecthanzzimmer"),
		Prefix: aws.String("batmanSoundtrack/"),
	})
	//Dune sountracks storage
	duneSoundrack, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String("petprojecthanzzimmer"),
		Prefix: aws.String("duneSoundtrack/"),
	})
	//Inception sountracks storage
	inceptionSoundrack, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String("petprojecthanzzimmer"),
		Prefix: aws.String("inceptionSoundtrack/"),
	})
	//PiratesOfTheCaribbean sountracks storage
	piratesOfTheCaribbeanSoundrack, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String("petprojecthanzzimmer"),
		Prefix: aws.String("piratesSoundtrack/"),
	})

	if err != nil {
		log.Fatal(err)
	}
	//Background Images
	var imageURLs []string
	// Генерируем URL для каждого объекта и выводим его
	for _, object := range backgoundimages.Contents {
		url := fmt.Sprintf("https://storage.yandexcloud.net/petprojecthanzzimmer/%s", aws.ToString(object.Key))
		imageURLs = append(imageURLs, url)
	}

	//Interstellar sountracks
	var interstellarSoundrackUrls []string
	for _, object := range interstellarSoundtrack.Contents {
		url := fmt.Sprintf("https://storage.yandexcloud.net/petprojecthanzzimmer/%s", aws.ToString(object.Key))
		interstellarSoundrackUrls = append(interstellarSoundrackUrls, url)
	}

	//Batman sountracks
	var batmanSoundrackUrls []string
	for _, object := range batmanSoundtrack.Contents {
		url := fmt.Sprintf("https://storage.yandexcloud.net/petprojecthanzzimmer/%s", aws.ToString(object.Key))
		batmanSoundrackUrls = append(batmanSoundrackUrls, url)
	}

	//Dune sountracks
	var duneSoundrackUrls []string
	for _, object := range duneSoundrack.Contents {
		url := fmt.Sprintf("https://storage.yandexcloud.net/petprojecthanzzimmer/%s", aws.ToString(object.Key))
		duneSoundrackUrls = append(duneSoundrackUrls, url)
	}

	//Inception sountracks
	var inceptionSoundrackUrls []string
	for _, object := range inceptionSoundrack.Contents {
		url := fmt.Sprintf("https://storage.yandexcloud.net/petprojecthanzzimmer/%s", aws.ToString(object.Key))
		inceptionSoundrackUrls = append(inceptionSoundrackUrls, url)
	}

	//PiratesOfTheCaribbeanSoundrack sountracks
	var piratesOfTheCaribbeanSoundrackUrls []string
	for _, object := range piratesOfTheCaribbeanSoundrack.Contents {
		url := fmt.Sprintf("https://storage.yandexcloud.net/petprojecthanzzimmer/%s", aws.ToString(object.Key))
		piratesOfTheCaribbeanSoundrackUrls = append(piratesOfTheCaribbeanSoundrackUrls, url)
	}
	return imageURLs[1:], interstellarSoundrackUrls[1:], batmanSoundrackUrls[1:], duneSoundrackUrls[1:], inceptionSoundrackUrls[1:], piratesOfTheCaribbeanSoundrackUrls[1:]

}
