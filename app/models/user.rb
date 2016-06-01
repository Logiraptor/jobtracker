class User < ActiveRecord::Base
  # Include default devise modules. Others available are:
  # :confirmable, :lockable, :timeoutable and :omniauthable
  devise :database_authenticatable, :registerable,
         :recoverable, :rememberable, :trackable, :validatable

  belongs_to :address
  validates :address, presence: true

  before_validation(on: :create) do
  	if address.nil?
  		self.address = Address.new
  	end
  end

  def full_name
  	[prefix, first_name, middle_initial, last_name, suffix].compact.join(" ")
  end
end
