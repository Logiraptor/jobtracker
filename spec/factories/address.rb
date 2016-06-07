FactoryGirl.define do
	factory :address do
		address_line_1 {Faker::Address.street_address}
		address_line_2 {Faker::Address.secondary_address}
		
		city {Faker::Address.city}
		state {Faker::Address.state}
		country {Faker::Address.country}
		zip_code {Faker::Address.zip_code}
	end
end