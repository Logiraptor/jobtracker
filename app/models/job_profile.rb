class JobProfile < ActiveRecord::Base
  enum salary_type: [:hourly, :monthly, :yearly]
  belongs_to :user
  belongs_to :address
  has_many :job_profile_duties
  has_many :job_profile_accomplishments
  has_many :job_profile_skills

  def human_start_date
    return '' if start_date.blank?
    start_date.strftime('%b %Y')
  end

  def human_end_date
    return 'present' if end_date.blank?
    end_date.strftime('%b %Y')
  end

  def duties
  	job_profile_duties.order(index: :asc).map(&:name)
  end
  def accomplishments
  	job_profile_accomplishments.order(index: :asc).map(&:name)
  end
  def skills
  	job_profile_skills.order(index: :asc).map(&:name)
  end
end
