require 'rails_helper'

describe JobProfile do
  describe 'human_start_date' do
    it 'is blank for nil start date' do
      profile = JobProfile.new(start_date: nil)
      expect(profile.human_start_date).to eq ''
    end
    it 'shows the month and year' do
      profile = JobProfile.new(start_date: Date.new(2015, 03, 06))
      expect(profile.human_start_date).to eq 'Mar 2015'
    end
  end
  describe 'human_end_date' do
    it 'is "present" for nil end date' do
      profile = JobProfile.new(end_date: nil)
      expect(profile.human_end_date).to eq 'present'
    end
    it 'shows the month and year' do
      profile = JobProfile.new(end_date: Date.new(2015, 03, 06))
      expect(profile.human_end_date).to eq 'Mar 2015'
    end
  end
  describe 'relationships' do
    before do
      duties = 5.downto(1).map do |i|
        JobProfileDuty.new(name: "Duty #{i}", index: i)
      end
      skills = 5.downto(1).map do |i|
        JobProfileSkill.new(name: "Skill #{i}", index: i)
      end
      accomplishments = 5.downto(1).map do |i|
        JobProfileAccomplishment.new(name: "Accomplishment #{i}", index: i)
      end
      @profile = create(:job_profile, job_profile_duties: duties,
                                      job_profile_skills: skills,
                                      job_profile_accomplishments: accomplishments)
    end
    describe 'duties' do
      it 'returns the duty names in order' do
        expect(@profile.duties).to eq [
          'Duty 1',
          'Duty 2',
          'Duty 3',
          'Duty 4',
          'Duty 5',
        ]
      end
    end
    describe 'skills' do
      it 'returns the skill names in order' do
        expect(@profile.skills).to eq [
          'Skill 1',
          'Skill 2',
          'Skill 3',
          'Skill 4',
          'Skill 5',
        ]
      end
    end
    describe 'accomplishments' do
      it 'returns the accomplishment names in order' do
        expect(@profile.accomplishments).to eq [
          'Accomplishment 1',
          'Accomplishment 2',
          'Accomplishment 3',
          'Accomplishment 4',
          'Accomplishment 5',
        ]
      end
    end
  end
end