
Dir.glob(File.expand_path("../factories/*.rb", __FILE__)).each do |file|
  require file
end

include FactoryGirl::Syntax::Methods

user = create(:user, email: 'test@gmail.com', password: 'password')

5.times do
	create(:job_profile_with_notes, user: user)
end