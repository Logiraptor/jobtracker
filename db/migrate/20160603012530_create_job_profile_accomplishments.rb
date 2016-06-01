class CreateJobProfileAccomplishments < ActiveRecord::Migration
  def change
    create_table :job_profile_accomplishments do |t|
      t.string :name
      t.belongs_to :job_profile, index: true, foreign_key: true
      t.integer :index

      t.timestamps null: false
    end
  end
end
