# encoding: utf-8

##
# This file is auto-generated. DO NOT EDIT!
#
require 'protobuf'
require 'protobuf/rpc/service'


##
# Imports
#
require 'user/v1/user.pb'

module User
  module V1
    ::Protobuf::Optionable.inject(self) { ::Google::Protobuf::FileOptions }

    ##
    # Message Classes
    #
    class GetUserRequest < ::Protobuf::Message; end
    class GetUserResponse < ::Protobuf::Message; end


    ##
    # File Options
    #
    set_option :go_package, "github.com/anhbkpro/videos/2022/07/23/gen/go/user/v1;userpb"


    ##
    # Message Fields
    #
    class GetUserRequest
      optional :string, :uuid, 1
    end

    class GetUserResponse
      optional ::User::V1::User, :user, 1
    end


    ##
    # Service Classes
    #
    class UserService < ::Protobuf::Rpc::Service
      rpc :get_user, ::User::V1::GetUserRequest, ::User::V1::GetUserResponse
    end

  end

end

