class AddAccountSettingsToUser < ActiveRecord::Migration
  def change
    change_table :users do |t|
      t.string :prefix
      t.string :suffix
      t.string :middle_initial
      t.string :phone_number
      t.string :website_url
      t.string :address_line_1
      t.string :address_line_2
      t.string :city
      t.string :state
      t.string :country
      t.string :zip_code
    end
  end
end
