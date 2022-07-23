# encoding: utf-8

##
# This file is auto-generated. DO NOT EDIT!
#
require 'protobuf'

module User
  module V1
    ::Protobuf::Optionable.inject(self) { ::Google::Protobuf::FileOptions }

    ##
    # Enum Classes
    #
    class MaritalStatus < ::Protobuf::Enum
      define :MARITAL_STATUS_UNSPECIFIED, 0
      define :MARITAL_STATUS_SINGLE, 1
      define :MARITAL_STATUS_MARRIED, 2
    end


    ##
    # Message Classes
    #
    class User < ::Protobuf::Message; end
    class Address < ::Protobuf::Message; end


    ##
    # File Options
    #
    set_option :go_package, "github.com/anhbkpro/videos/2022/07/23/gen/go/user/v1;userpb"


    ##
    # Message Fields
    #
    class User
      optional :string, :uuid, 1
      optional :string, :full_name, 2
      optional :int64, :birth_year, 3
      optional :uint32, :salary, 4
      repeated ::User::V1::Address, :addresses, 5
      optional ::User::V1::MaritalStatus, :marital_status, 6
    end

    class Address
      optional :string, :street, 1
      optional :string, :city, 2
    end

  end

end

