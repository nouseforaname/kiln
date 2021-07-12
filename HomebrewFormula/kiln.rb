# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class Kiln < Formula
  desc ""
  homepage ""
  version "0.60.2"
  bottle :unneeded

  on_macos do
    if Hardware::CPU.intel?
      url "https://github.com/pivotal-cf/kiln/releases/download/0.60.2/kiln-darwin-0.60.2.tar.gz"
      sha256 "e4ba72335f8b6eb26a9cbf86ed3e53e40b5ad3528fd73476958d8fb0867b7750"
    end
  end

  on_linux do
    if Hardware::CPU.intel?
      url "https://github.com/pivotal-cf/kiln/releases/download/0.60.2/kiln-linux-0.60.2.tar.gz"
      sha256 "5fdb609f06ece9f7e7fba2e1d107d9dcc2ab1c71dda0fcd86b03f57952b650fa"
    end
  end

  def install
    bin.install "kiln"
  end

  test do
    system "#{bin}/kiln --version"
  end
end
