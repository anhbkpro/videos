# encoding: utf-8

##
# This file is auto-generated. DO NOT EDIT!
#
require 'protobuf'
require 'protobuf/rpc/service'

module Wearable
  module V1
    ::Protobuf::Optionable.inject(self) { ::Google::Protobuf::FileOptions }

    ##
    # Message Classes
    #
    class BeatsPerMinuteRequest < ::Protobuf::Message; end
    class BeatsPerMinuteResponse < ::Protobuf::Message; end
    class ConsumeBeatsPerMinuteRequest < ::Protobuf::Message; end
    class ConsumeBeatsPerMinuteResponse < ::Protobuf::Message; end
    class CalculateBeatsPerMinuteRequest < ::Protobuf::Message; end
    class CalculateBeatsPerMinuteResponse < ::Protobuf::Message; end


    ##
    # File Options
    #
    set_option :go_package, "github.com/anhbkpro/videos/2022/07/23/gen/go/wearable/v1;wearablepb"


    ##
    # Message Fields
    #
    class BeatsPerMinuteRequest
      optional :string, :uuid, 1
    end

    class BeatsPerMinuteResponse
      optional :uint32, :value, 1
      optional :uint32, :minute, 2
    end

    class ConsumeBeatsPerMinuteRequest
      optional :string, :uuid, 1
      optional :uint32, :value, 2
      optional :uint32, :minute, 3
    end

    class ConsumeBeatsPerMinuteResponse
      optional :uint32, :total, 1
    end

    class CalculateBeatsPerMinuteRequest
      optional :string, :uuid, 1
      optional :uint32, :value, 2
      optional :uint32, :minute, 3
    end

    class CalculateBeatsPerMinuteResponse
      optional :float, :average, 1
    end


    ##
    # Service Classes
    #
    class WearableService < ::Protobuf::Rpc::Service
      rpc :beats_per_minute, ::Wearable::V1::BeatsPerMinuteRequest, ::Wearable::V1::BeatsPerMinuteResponse
      rpc :consume_beats_per_minute, ::Wearable::V1::ConsumeBeatsPerMinuteRequest, ::Wearable::V1::ConsumeBeatsPerMinuteResponse
      rpc :calculate_beats_per_minute, ::Wearable::V1::CalculateBeatsPerMinuteRequest, ::Wearable::V1::CalculateBeatsPerMinuteResponse
    end

  end

end

