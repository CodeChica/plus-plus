class CreateSparkles < ActiveRecord::Migration[6.1]
  def change
    create_table :sparkles do |t|
      t.string :reason
      t.references :sparkler, index: true, null: false
      t.references :sparklee, index: true, null: false

      t.timestamps
    end
  end
end
