class CreateUsers < ActiveRecord::Migration[6.1]
  def change
    create_table :users do |t|
      t.string :handle, null: false
      t.timestamps
    end
    add_index :users, :handle, unique: true
  end
end
