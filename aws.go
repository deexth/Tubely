package main

// func (cfg *apiConfig) dbVideoToSignedVideo(video database.Video) (database.Video, error) {
// 	// 1. Check if VideoURL is nil
// 	if video.VideoURL == nil {
// 		return video, nil
// 	}
//
// 	// 2. Split and check if we actually got two parts (bucket and key)
// 	d := strings.Split(*video.VideoURL, ",")
// 	if len(d) < 2 {
// 		return video, nil
// 	}
// 	presinedURL, err := generatePresignedURL(cfg.s3Client, d[0], d[1], time.Second*5)
// 	if err != nil {
// 		return database.Video{}, err
// 	}
//
// 	video.VideoURL = &presinedURL
//
// 	return video, nil
// }
//
// func generatePresignedURL(s3Client *s3.Client, bucket, key string, expireTime time.Duration) (string, error) {
// 	presignedClient := s3.NewPresignClient(s3Client)
//
// 	presignedReq, err := presignedClient.PresignGetObject(context.Background(), &s3.GetObjectInput{
// 		Bucket: &bucket,
// 		Key:    &key,
// 	}, s3.WithPresignExpires(expireTime))
// 	if err != nil {
// 		return "", fmt.Errorf("issue signing URL: %v", err)
// 	}
//
// 	return presignedReq.URL, nil
// }
