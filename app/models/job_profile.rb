class JobProfile < ActiveRecord::Base
  enum salary_type: [:hourly, :monthly, :yearly]
  belongs_to :address
  has_many :job_profile_duties
  has_many :job_profile_accomplishments
  has_many :job_profile_skills

  def duties
  	job_profile_duties.map(&:name)
  end
  def accomplishments
  	job_profile_accomplishments.map(&:name)
  end
  def skills
  	job_profile_skills.map(&:name)
  end
end
