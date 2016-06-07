require 'rails_helper'

describe 'generating a resume' do
  before do
  	@user = create(:user)
  	@job_history = (0..5).map {|i| create(:job_profile_with_notes, user: @user) }

  	login_as(@user)
  	visit '/'
  	click_on 'Resumes'
  	click_on 'Generate'
  end

  it 'contains the users contact info' do
    expect(page).to have_content @user.full_name
    expect(page).to have_content @user.address.address_line_1
    expect(page).to have_content @user.address.address_line_2
    expect(page).to have_content @user.address.city
    expect(page).to have_content @user.address.state
  	expect(page).to have_content @user.address.zip_code
  end

  it 'contains the users job history' do
    @job_history.each do |job|
      expect(page).to have_content job.title
      expect(page).to have_content job.employer_name
      expect(page).to have_content job.human_start_date
      expect(page).to have_content job.human_end_date

      job.skills.each do |skill|
        expect(page).to have_content skill
      end

      job.duties.each do |duty|
        expect(page).to have_content duty
      end

      job.accomplishments.each do |accomplishment|
        expect(page).to have_content accomplishment
      end
    end
  end
end