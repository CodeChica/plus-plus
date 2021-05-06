class Sparkle < ApplicationRecord
  belongs_to :sparkler, class_name: 'User'
  belongs_to :sparklee, class_name: 'User'

  validates :reason, length: { maximum: 255 }, allow_blank: true
  validates :sparkler, :sparklee, presence: true

  scope :recent, ->() { order(created_at: :desc) }
end
