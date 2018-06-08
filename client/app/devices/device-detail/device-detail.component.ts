import { Component, OnInit } from "@angular/core";
import { PageRoute, RouterExtensions } from "nativescript-angular/router";
import { switchMap } from "rxjs/operators";

import { DeviceInfo } from "../../shared/device.model";
import { DevicesService } from "../../shared/devices-service";
import { ActivatedRoute } from "@angular/router";
import { Observable } from "tns-core-modules/ui/page/page";

/* ***********************************************************
* This is the item details component in the master-detail structure.
* This component retrieves the passed parameter from the master list component,
* finds the data item by this parameter and displays the detailed data item information.
*************************************************************/
@Component({
    moduleId: module.id,
    templateUrl: "./device-detail.component.html"
})
export class DeviceDetailComponent implements OnInit {
    private _device: DeviceInfo;
    private _user: string;
    constructor(
        private route: ActivatedRoute,
        private _devicesService: DevicesService,
        private _pageRoute: PageRoute,
        private _routerExtensions: RouterExtensions
    ) { }

    /* ***********************************************************
    * Use the "ngOnInit" handler to get the data item id parameter passed through navigation.
    * Get the data item details from the data service using this id and assign it to the
    * private property that holds it inside the component.
    *************************************************************/

    ngOnInit(): void {
        /* ***********************************************************
        * Learn more about how to get navigation parameters in this documentation article:
        * http://docs.nativescript.org/angular/core-concepts/angular-navigation.html#passing-parameter
        *************************************************************/
        this.route.queryParams
            .subscribe(params => {
                console.log("PARAMS!!!!!", params);
                this._user = params.user;
                const deviceId = params.deviceId;
                this._devicesService.getDeviceInfo(deviceId, this._user)
                    .then(deviceInfo =>{
                        this._device = deviceInfo;
                        console.log("@@@@@@ SET DEVICE TO : ", this._device);
                    });
            });

        // this._pageRoute.activatedRoute
        //     .pipe(switchMap((activatedRoute) => activatedRoute.params))
        //     .forEach((params) => {
        //         console.log("######", params);
        //         const deviceId = params.id;
        //         this._user = params.user;

        //         console.log("@@@@@@@@@@@@@@@", deviceId, this._user);

        //         this._devicesService.getDeviceInfo(deviceId, this._user)
        //             .then(deviceInfo => this._device = deviceInfo);
        //     });
    }

    get device(): DeviceInfo {
        console.log("@@@@@@@@@@@!!!!!!!!!!!!!! getting device", this._device);
        return this._device;
    }

    /* ***********************************************************
    * The back button is essential for a master-detail feature.
    *************************************************************/
    onBackButtonTap(): void {
        this._routerExtensions.backToPreviousPage();
    }
}
