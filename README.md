# devices-rack-organizer

The application's purpose is to allow tracking of the devices we have in the office. When device is taken from the device drawer, you'll have to identify yourself and scan the QR code of the device. Once you do this, information that you've taken it will be persisted in the database.
When you want to return it, just use the return button, so you'll now that it is available.
Application can show which are the currently avaialble devices.

The application's client side is mobile application written in NativeScript. The server side is Serverless Go lang application.

## Video
Video showing the app can be found here: https://www.screencast.com/t/pAbLWC7TeoQR

## Artifacts
You can find .apk [here](https://github.com/zahhak/devices-rack-organizer/releases/download/v0.1.0/app-release.apk).


## Client side
The client side is a NativeScript application. It uses Angular 6 and webpack 4.0.0 to improve the performance. Application has a login screen. After entering the username, you are forwarded to the main screen. At this point there are three actions that you can do:
* Get device - just use the `GET` button and scan the QR code of the device you want to use.
* Return device - after tapping on the `RETURN` button, you can scan the code on the device that you want to return in the drawer.
* List devices - tap on the `LIST` button and you'll get full list of all devices, which are free and who has taken the others.

When you check the List of devices, you can select a specific one and you'll see the full history for this device - when it has been taken, returned, etc.

### Tools and technologies
* NativeScript Sidekick 1.10.2-v.2018.6.5.2
* NativeScript CLI 4.1.0
* Visual Studio Code 1.23.1
* Webpack 4.0.0
* Android emulators
* NativeScript 4.1.0
* TypeScript 2.7.2
* Angular 6.0.0

### Development experience
During development we've identified several issues. They are described in [this document](https://progresssoftware-my.sharepoint.com/:w:/g/personal/vladimirov_progress_com/ETCsVXih0G9LgaJ6sVy90okB7QI0PCElf_yZizdaIYXbEw?e=b0jJmq).

Android build is quite slow in the latest version. Also the LiveSync experience is slow when changing .html files.

### Times:
- App Startup time on LG G4 with snapshot, AOT and uglify +1s394ms
- Startup time with debug builds and no optimization +4s116ms

## Server side
Backend service consist of the following components working together:

AWS API Gateway is accepting the HTTP call and pass it to AWS Lambda. Lambda will call the AWS DynamoDB. Furthermore, in the JSON response, S3 URLs for image and history will be returned.

### Tools:

* Visual Studio Code
* Golang
* Third-party dependencies:
* Terraform
* Serveless
* AWS Lambda
* AWS API Gateway
* AWS Dynamodb
* AWS S3

### Development experience
For our App (Rack Organizer) AWS Lambda + AWS API Gateway + AWS DynamoDB + AWS S3 is a good fit as it allows us to quickly create a RESTful API. The notion of different Stages supported by the AWS technologies support our development workflow. Golang has a very good Visual Studio Code Support which speed up the development part.  For this simple App Terraform had no problems creating the Infrastructure layout.
