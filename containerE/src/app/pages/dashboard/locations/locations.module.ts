import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { LocationsComponent } from '@pages/dashboard/locations/locations.component';
import { PageTemplateModule } from '@shared/components/page-template/page-template.module';



@NgModule({
  declarations: [
    LocationsComponent
  ],
  imports: [
    CommonModule,
    PageTemplateModule
  ]
})
export class LocationsModule { }
