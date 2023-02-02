output "region" {
  value       = var.region
  description = "gcp/gke region"
}

output "project_id" {
  value       = var.project_id
  description = "gcp/gke project id"
}

output "gke_cluster_name" {
  value       = google_container_cluster.primary.name
  description = "gcp/gke cluster name"
}
