# Shopify to Buchhaltungsbutler CSV Parser

## Overview

This tool serves as a bridge between your Shopify store and the Buchhaltungsbutler accounting software. It takes CSV reports from Shopify, converts them into a format compatible with Buchhaltungsbutler, and prepares them for manual upload into Buchhaltungsbutler.

## Prerequisites

- An active Shopify store
- A Buchhaltungsbutler account

## Installation

1. Download the latest executable from the release page: [GitHub Releases](https://github.com/DerbeDotDev/butler-csv/releases)

2. Give the downloaded executable permission to run:
    ```bash
    chmod +x butler-csv
    ```

## Usage

1. Download both the Transaction Report and the Payout Report for the desired time period from your Shopify store as CSV files.

2. Open your terminal and navigate to the location of the executable.

3. Run the parser using the following command:
    ```bash
    ./butler-csv --input=/path/to/input --output=/path/to/output
    ```
   Replace `/path/to/input` and `/path/to/output` with the actual paths where the input CSV files are located and where the output CSV files should be saved.

4. After successful execution, you'll find the CSV files prepared for Buchhaltungsbutler at the specified output path.

5. Manually upload these files into your Buchhaltungsbutler account.

## Notes

This tool is an automated solution for transferring data from Shopify to Buchhaltungsbutler but still requires manual upload of the generated CSV files into Buchhaltungsbutler.
