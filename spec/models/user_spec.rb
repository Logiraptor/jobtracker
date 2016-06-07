require 'rails_helper'

describe User do
	describe '#address' do
		it 'defaults to an empty address' do
			user = create(:user, address_id: nil)
			expect(user.address.persisted?)
		end
	end

	describe '#full_name' do
		it 'returns all present parts joined with space' do
			user = build(:user,
				prefix: 'prefix',
				first_name: 'first_name',
				middle_initial: 'middle_initial',
				last_name: 'last_name',
				suffix: 'suffix'
			)
			expect(user.full_name).to eq 'prefix first_name middle_initial last_name suffix'
		end
		it 'excludes nil elements' do
			user = build(:user,
				prefix: nil,
				first_name: nil,
				middle_initial: 'middle_initial',
				last_name: nil,
				suffix: 'suffix'
			)
			expect(user.full_name).to eq 'middle_initial suffix'
		end
	end
end