class UsersController < ApplicationController
  before_action :authenticate_user!
  
  def show
  end

  def edit
  end

  def update
    current_user.update!(**user_params)
    redirect_to current_user
  end

  private

  def user_params
    params.require(:user).permit(
      :first_name,
      :last_name,
      :prefix,
      :suffix,
      :middle_initial,
      :phone_number,
      :website_url,
      :address_line_1,
      :address_line_2,
      :city,
      :state,
      :country,
      :zip_code
    ).symbolize_keys
  end
end
