class ResumesController < ApplicationController
	before_action :authenticate_user!

	def index
		@user = current_user
		@job_profiles = @user.job_profiles.includes(
			:job_profile_duties,
			:job_profile_accomplishments,
			:job_profile_skills
		)

		respond_to do |format|
	      format.html do
	      	render template: 'resumes/index.pdf.erb'
	      end
	      format.pdf do
	        render pdf: "resume",
	        	   layout: 'pdf.html.erb',
	        	   show_as_html: params.key?('debug')
	      end
	    end
	end

	def create
		redirect_to resumes_path, format: 'pdf'
	end
end
