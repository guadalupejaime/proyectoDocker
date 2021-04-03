import { Component, Input, OnInit } from '@angular/core';

@Component({
  selector: 'app-page-template',
  templateUrl: './page-template.component.html',
  styleUrls: ['./page-template.component.scss']
})
export class PageTemplateComponent implements OnInit {
  @Input() status: 'load' | 'ok' | 'bad' = 'load';
  @Input() title: string;

  constructor() { }

  ngOnInit(): void {
  }

}
