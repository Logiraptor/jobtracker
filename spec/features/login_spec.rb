require 'rails_helper'


describe 'home page' do
	before do
		visit '/'
		create(:user, first_name: 'Fred', email: 'test@example.com', password: 'password')
	end

	context 'while not logged in' do
		it 'shows a login form' do
			fill_in 'Email', with: 'test@example.com'
			fill_in 'Password', with: 'password'
			click_on 'Log in'

			expect(page).to have_content('Fred')
		end
	end
end