gcloud functions deploy prod-stars-cf \
  --region=asia-east1 \
  --service-account sa-transcoder@jkface.iam.gserviceaccount.com \
  --entry-point "Entry" \
  --runtime go116 \
  --memory 128 \
  --trigger-event google.storage.object.finalize \
  --trigger-resource prod-stars-public