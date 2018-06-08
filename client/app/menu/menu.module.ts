import { NgModule, NO_ERRORS_SCHEMA } from "@angular/core";
import { Routes } from "@angular/router";
import { MenuComponent } from "~/menu/menu.component";
import { NativeScriptRouterModule } from "nativescript-angular/router";
import { NativeScriptFormsModule } from "nativescript-angular/forms";
import { DevicesService } from "~/shared/devices-service";

const routes: Routes = [
    { path: "", component: MenuComponent }
];

@NgModule({
    imports: [
        NativeScriptRouterModule.forChild(routes),
        NativeScriptFormsModule
    ],
    declarations: [
        MenuComponent
    ],
    providers: [
        DevicesService
    ],
    bootstrap: [MenuComponent],
    schemas: [
        NO_ERRORS_SCHEMA
    ]
})
export class MenuModule { }
