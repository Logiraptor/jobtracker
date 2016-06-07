class JobProfilesController < ApplicationController
  include Address::Params
  before_action :set_job_profile, only: [:show, :edit, :update, :destroy]

  # GET /job_profiles
  # GET /job_profiles.json
  def index
    @job_profiles = current_user.job_profiles
  end

  # GET /job_profiles/1
  # GET /job_profiles/1.json
  def show
  end

  # GET /job_profiles/new
  def new
    @job_profile = JobProfile.new
  end

  # GET /job_profiles/1/edit
  def edit
  end

  # POST /job_profiles
  # POST /job_profiles.json
  def create
    @job_profile = JobProfile.new(job_profile_params)
    @job_profile.user = current_user
    @job_profile.address =  Address.new(address_params)
    @job_profile.job_profile_duties = multi_form_params(:duties).map do |duty|
      JobProfileDuty.new(name: duty)
    end
    @job_profile.job_profile_accomplishments = multi_form_params(:accomplishments).map do |accomplishment|
      JobProfileAccomplishment.new(name: accomplishment)
    end
    @job_profile.job_profile_skills = multi_form_params(:skills).map do |skill|
      JobProfileSkill.new(name: skill)
    end

    respond_to do |format|
      if @job_profile.save
        format.html { redirect_to @job_profile, notice: 'Job Profile was successfully created.' }
        format.json { render :show, status: :created, location: @job_profile }
      else
        format.html { render :new }
        format.json { render json: @job_profile.errors, status: :unprocessable_entity }
      end
    end
  end

  # PATCH/PUT /job_profiles/1
  # PATCH/PUT /job_profiles/1.json
  def update
    @job_profile.address =  Address.new(address_params)
    @job_profile.job_profile_duties = multi_form_params(:duties).map do |duty|
      JobProfileDuty.new(name: duty)
    end
    @job_profile.job_profile_accomplishments = multi_form_params(:accomplishments).map do |accomplishment|
      JobProfileAccomplishment.new(name: accomplishment)
    end
    @job_profile.job_profile_skills = multi_form_params(:skills).map do |skill|
      JobProfileSkill.new(name: skill)
    end

    respond_to do |format|
      if @job_profile.save && @job_profile.update(job_profile_params)
        format.html { redirect_to @job_profile, notice: 'Job Profile was successfully updated.' }
        format.json { render :show, status: :ok, location: @job_profile }
      else
        format.html { render :edit }
        format.json { render json: @job_profile.errors, status: :unprocessable_entity }
      end
    end
  end

  # DELETE /job_profiles/1
  # DELETE /job_profiles/1.json
  def destroy
    @job_profile.destroy
    respond_to do |format|
      format.html { redirect_to job_profiles_url, notice: 'Job Profile was successfully destroyed.' }
      format.json { head :no_content }
    end
  end

  private
    # Use callbacks to share common setup or constraints between actions.
    def set_job_profile
      @job_profile = JobProfile.find(params[:id])
    end

    # Never trust parameters from the scary internet, only allow the white list through.
    def job_profile_params
      params.require(:job_profile).permit(
        :employer_name,
        :title,
        :start_date,
        :end_date,
        :salary,
        :salary_type,
        :average_weekly_hours,
        :supervisor_name,
        :supervisor_phone
      )
    end

    def multi_form_params(name)
      params.require(name)[:name]
    end
end
