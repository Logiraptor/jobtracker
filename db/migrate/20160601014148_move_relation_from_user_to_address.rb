class MoveRelationFromUserToAddress < ActiveRecord::Migration
  def change
  	remove_reference :users, :address

  	change_table :users do |t|
  		t.belongs_to :address
  	end
  end
end
