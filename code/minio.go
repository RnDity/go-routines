func uploadFilesToMinio(fs []*multipart.FileHeader, client *minio.Client, customerID string, userID string, artifactIDs []string) error {
	var wg sync.WaitGroup
	errch := make(chan error, len(fs))
	for i, file := range fs {
		go func(f *multipart.FileHeader, artifactID string) {
			defer wg.Done()
			reader, err := f.Open()
			if err != nil {
				errch <- errors.Wrap(err, "failed uploading file")
				return
			}
			defer reader.Close()
			if _, err = client.PutObject(
				customerID,
				filepath.Join("artifacts", userID, artifactID),
				reader,
				f.Size,
				minio.PutObjectOptions{ContentType: f.Header.Get("Content-Type")},
			); err != nil {
				errch <- errors.Wrap(err, "failed putting file to storage")
				return
			}
		}(file, artifactIDs[i])
	}
	wg.Add(len(fs))
	wg.Wait()
	select {
	case err, ok := <-errch:
		if ok {
			return err
		}
		return nil
	default:
		return nil
	}
}
