class AddFirstAndLastNameToUser < ActiveRecord::Migration
  def change
    change_table :users do |t|
      t.string :first_name, default: '', null: false
      t.string :last_name, default: '', null: false
    end
  end
end
