module Rack
  module Session
    class MessagePack < Cookie::Base64
      def encode(obj)
        super(::MessagePack.pack(obj))
      end

      def decode(str)
        return unless str
        ::MessagePack.unpack(super(str)) rescue nil
      end
    end
  end
end

