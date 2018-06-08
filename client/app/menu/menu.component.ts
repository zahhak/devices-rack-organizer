import { Component } from "@angular/core";
import { ActivatedRoute } from "@angular/router";
import { PageRoute, RouterExtensions } from "nativescript-angular/router";
import { switchMap } from "rxjs/operators";
import { BarcodeScanner } from "nativescript-barcodescanner";

import { HttpClient } from "~/shared/http-client";
import { DevicesService } from "~/shared/devices-service";
import * as dialogs from "ui/dialogs";

@Component({
    moduleId: module.id,
    templateUrl: "./menu.component.html"
})

export class MenuComponent {

    public user: string;
    constructor(
        private route: ActivatedRoute,
        private _routerExtensions: RouterExtensions,
        private _devicesService: DevicesService) { }

    public updateDevice(): Promise<any> {
        const scanner = new BarcodeScanner();
        let deviceAction = "";
        let deviceId = "";
        return scanner
            .scan({
                cancelLabel: "Stop scanning",
                message: "Scan the code on the back of the device",
                preferFrontCamera: false,
                showFlipCameraButton: true
            })
            .then(result => {
                deviceId = result.text;
                return this._devicesService.updateDeviceInfo(deviceId, this.user)
            })
            .then(deviceStatus => {
                deviceAction = deviceStatus.deviceState;
                return this._routerExtensions.navigate([""], {
                    clearHistory: true,
                    animated: true,
                    transition: {
                        name: "slideBottom",
                        duration: 200,
                        curve: "ease"
                    }
                });
            })
            .then(() => dialogs.alert({
                title: "Success",
                message: `You have successfully ${deviceAction.toLowerCase()} device ${deviceId}`,
                okButtonText: "Got it"
            }));
    }

    ngOnInit() {
        this.route.queryParams
            .subscribe(params => {
                this.user = params.user;
            });
    }

    public onListTapped() {
        return this._routerExtensions.navigate(["/devices"], {
            clearHistory: true,
            animated: true,
            transition: {
                name: "slideBottom",
                duration: 200,
                curve: "ease"
            },
            queryParams: {
                user: this.user
            }
        });

        // this._devicesService.getAllDevices(this.user)
        //     .then(res => alert("Result is " + JSON.stringify(res, null, 2)))
        //     .catch(err => alert("error while getting devices: " + err));

        console.log("@@@@@@@@@@@@ LS");
        alert("TAPPED!!!!!");
    }

}
