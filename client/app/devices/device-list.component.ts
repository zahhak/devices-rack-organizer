import { Component, OnDestroy, OnInit } from "@angular/core";
import { ObservableArray } from "data/observable-array";
import { RouterExtensions } from "nativescript-angular/router";
import { ListViewEventData } from "nativescript-ui-listview";
import { Subscription } from "rxjs";
import { finalize } from "rxjs/operators";

import { DeviceInfo } from "../shared/device.model";
import { DevicesService } from "../shared/devices-service";
import { ActivatedRoute } from "@angular/router";

@Component({
    moduleId: module.id,
    templateUrl: "./devices-list.component.html",
    styleUrls: ["./car-list.component.scss"]
})
export class DevicesListComponent implements OnInit {
    private _isLoading: boolean = false;
    private _devices: DeviceInfo[] = [];
    private _dataSubscription: Subscription;

    public user: string;

    constructor(
        private route: ActivatedRoute,
        private _devicesService: DevicesService,
        private _routerExtensions: RouterExtensions
    ) { }


    ngOnInit(): void {
        this._isLoading = true;
        this.route.queryParams
            .subscribe(params => {
                this.user = params.user;
                this._devicesService.getAllDevices(this.user)
                    .then(res => {
                        this._devices = res;
                        this._isLoading = false;
                    });
            });
    }

    get devices(): DeviceInfo[] {
        return this._devices;
    }

    get isLoading(): boolean {
        return this._isLoading;
    }

    onDeviceItemTap(args: ListViewEventData): void {
        const tappedDeviceItem = args.view.bindingContext;

        this._routerExtensions.navigate(["/devices/device-detail"],
            {
                animated: true,
                transition: {
                    name: "slide",
                    duration: 200,
                    curve: "ease"
                },
                queryParams: {
                    deviceId: tappedDeviceItem.deviceId,
                    user: this.user
                }
            });
    }
}
