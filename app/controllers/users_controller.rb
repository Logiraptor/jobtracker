class UsersController < ApplicationController
  include Address::Params
  before_action :authenticate_user!

  def update
    current_user.update!(**user_params)
    current_user.address.update!(**address_params)
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
      :website_url,
    ).symbolize_keys
  end
end
