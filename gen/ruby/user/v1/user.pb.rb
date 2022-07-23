# encoding: utf-8

##
# This file is auto-generated. DO NOT EDIT!
#
require 'protobuf'

module User
  module V1
    ::Protobuf::Optionable.inject(self) { ::Google::Protobuf::FileOptions }

    ##
    # Message Classes
    #
    class User < ::Protobuf::Message; end


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
    end

  end

end

