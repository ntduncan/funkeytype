class Funkeytype < Formula
  desc "Typing test in your terminal"
  homepage "https://github.com/your-username/funkeytype"
  url "https://github.com/your-username/funkeytype/archive/v1.0.0.tar.gz" # Replace with your actual release tarball
  sha256 "your_sha256_checksum_here" # Replace with the actual checksum of your release tarball
  license "MIT"

  depends_on "go" => :build

  def install
    system "go", "build", "-o", "funkeytype", "."
    bin.install "funkeytype"
  end

  test do
    # This is a basic test, you can expand it
    assert_match "funkeytype version", shell_output("#{bin}/funkeytype --version")
  end
end
