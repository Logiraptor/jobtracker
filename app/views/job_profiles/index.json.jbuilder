json.array!(@job_profiles) do |job_profile|
  json.extract! job_profile, :id
  json.url job_profile_url(job_profile, format: :json)
end
