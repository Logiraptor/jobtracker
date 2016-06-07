
namespace :dev do
	task reset: ['db:drop', 'db:create', 'db:migrate'] do
		require_relative '../../spec/support/example_data'
	end
end