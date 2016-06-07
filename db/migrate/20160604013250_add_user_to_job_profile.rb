class AddUserToJobProfile < ActiveRecord::Migration
  def change
  	change_table :job_profiles do |t|
  		t.belongs_to :user, foreign_key: true
  	end
  end
end
