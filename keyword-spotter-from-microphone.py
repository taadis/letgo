#!/usr/bin/env python3

# Real-time keyword spotting from a microphone with sherpa-onnx Python API
#
# Please refer to
# https://k2-fsa.github.io/sherpa/onnx/kws/pretrained_models/index.html
# to download pre-trained models

import argparse
import sys
from pathlib import Path

from typing import List

try:
    import sounddevice as sd
except ImportError:
    print("Please install sounddevice first. You can use")
    print()
    print("  pip install sounddevice")
    print()
    print("to install it")
    sys.exit(-1)

try:
    import sherpa_onnx
except ImportError:
    print("Please install sherpa_onnx first. You need it.")
    print()
    print(" pip install sherpa_onnx")
    print()
    print("to install it.")
    sys.exit(-1)

def main():
    devices = sd.query_devices()
    if len(devices) == 0:
        print("No microphone devices found")
        sys.exit(0)

    print(devices)
    default_input_device_idx = sd.default.device[0]
    print(f'Use default device: {devices[default_input_device_idx]["name"]}')

    keyword_spotter = sherpa_onnx.KeywordSpotter(
        tokens="tokens.txt",
        encoder="encoder-epoch-12-avg-2-chunk-16-left-64.onnx",
        decoder="decoder-epoch-12-avg-2-chunk-16-left-64.onnx",
        joiner="joiner-epoch-99-avg-1-chunk-16-left-64.onnx",
        num_threads=4,
        max_active_paths=3,
        keywords_file="keywords.txt",
        keywords_score=2.2,
        keywords_threshold=0.25,
        num_trailing_blanks=1,
        provider="cpu",
    )

    print("Started! Please speak")

    idx = 0

    sample_rate = 16000
    samples_per_read = int(0.1 * sample_rate)  # 0.1 second = 100 ms
    stream = keyword_spotter.create_stream()
    with sd.InputStream(channels=1, dtype="float32", samplerate=sample_rate) as s:
        while True:
            samples, _ = s.read(samples_per_read)  # a blocking read
            samples = samples.reshape(-1)
            stream.accept_waveform(sample_rate, samples)
            while keyword_spotter.is_ready(stream):
                keyword_spotter.decode_stream(stream)
                result = keyword_spotter.get_result(stream)
                if result:
                    print(f"{idx}: {result }")
                    idx += 1
                    # Remember to reset stream right after detecting a keyword
                    keyword_spotter.reset_stream(stream)


if __name__ == "__main__":
    try:
        main()
    except KeyboardInterrupt:
        print("\nCaught Ctrl + C. Exiting")
