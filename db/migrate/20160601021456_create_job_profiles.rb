class CreateJobProfiles < ActiveRecord::Migration
  def change
    create_table :job_profiles do |t|
      t.string :title
      t.string :employer_name
      t.belongs_to :address, index: true, foreign_key: true
      
      t.datetime :start_date
      t.datetime :end_date
      t.decimal :salary
      t.integer :salary_type
      t.integer :average_weekly_hours

      t.string :supervisor_name
      t.string :supervisor_phone

      t.timestamps null: false
    end
  end
end
