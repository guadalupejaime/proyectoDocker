import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { PageTemplateComponent } from '@shared/components/page-template/page-template.component';
import { MyMaterialModule } from '@shared/my-material/my-material.module';



@NgModule({
  declarations: [
    PageTemplateComponent
  ],
  exports: [
    PageTemplateComponent
  ],
  imports: [
    CommonModule,
    MyMaterialModule
  ]
})
export class PageTemplateModule { }
