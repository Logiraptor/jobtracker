class Address < ActiveRecord::Base
	module Params
		def address_params
			params.require(:address).permit(
		      :phone_number,
		      :address_line_1,
		      :address_line_2,
		      :city,
		      :state,
		      :country,
		      :zip_code
		    ).symbolize_keys
		end
	end

end
