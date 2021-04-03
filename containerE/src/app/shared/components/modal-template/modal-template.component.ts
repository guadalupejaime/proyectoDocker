import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';

@Component({
  selector: 'app-modal-template',
  templateUrl: './modal-template.component.html',
  styleUrls: ['./modal-template.component.scss']
})
export class ModalTemplateComponent implements OnInit {
  @Input() status: 'load' | 'ok' | 'bad' = 'load';
  @Input() titulo = 'noTitle';
  @Input() widthModal = '450px';
  @Output() closeModal = new EventEmitter();

  constructor() { }

  ngOnInit(): void {
  }

}
