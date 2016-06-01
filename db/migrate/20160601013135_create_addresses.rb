class CreateAddresses < ActiveRecord::Migration
  def change
    create_table :addresses do |t|
      t.string :phone_number
      t.string :address_line_1
      t.string :address_line_2
      t.string :city
      t.string :state
      t.string :country
      t.string :zip_code

      t.timestamps null: false
    end

    change_table :users do |t|
      t.remove :phone_number
      t.remove :address_line_1
      t.remove :address_line_2
      t.remove :city
      t.remove :state
      t.remove :country
      t.remove :zip_code
    end

    add_reference :users, :address, index: true
  end
end
