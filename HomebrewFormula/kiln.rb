# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class Kiln < Formula
  desc ""
  homepage ""
  version "0.62.3-rc3"

  on_macos do
    if Hardware::CPU.intel?
      url "https://github.com/pivotal-cf/kiln/releases/download/0.62.3-rc3/kiln-darwin-0.62.3-rc3.tar.gz"
      sha256 "5f4ef5e4cc0ab14d056f721275ca0a704041790cfce54a57031cf82ed86949a7"

      def install
        bin.install "kiln"
      end
    end
  end

  on_linux do
    if Hardware::CPU.intel?
      url "https://github.com/pivotal-cf/kiln/releases/download/0.62.3-rc3/kiln-linux-0.62.3-rc3.tar.gz"
      sha256 "37a8221bfa942a1f3894a592fa3d8333e57ab40a7d7ad481b2188542abb87f3d"

      def install
        bin.install "kiln"
      end
    end
  end

  test do
    system "#{bin}/kiln --version"
  end
end
