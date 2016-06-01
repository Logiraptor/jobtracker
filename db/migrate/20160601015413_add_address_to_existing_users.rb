class AddAddressToExistingUsers < ActiveRecord::Migration
  def change
  	User.where(address_id: nil).each do |user|
  		user.address = Address.new
  		user.save!
  	end
  end
end
