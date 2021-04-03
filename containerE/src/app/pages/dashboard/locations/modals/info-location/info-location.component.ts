import { Component, Inject, OnInit } from '@angular/core';
import { MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { Location } from '@pages/dashboard/locations/locations.interface';
import { HttpService } from '@services/http.service';

@Component({
  selector: 'app-info-location',
  templateUrl: './info-location.component.html',
  styleUrls: ['./info-location.component.scss']
})
export class InfoLocationComponent implements OnInit {
  status = 'load';
  location: Location;
  characters = [];

  constructor(
    public sMatDialogRef: MatDialogRef<InfoLocationComponent>,
    @Inject(MAT_DIALOG_DATA) public data: { idLocation: string },
    private sHttpService: HttpService
  ) {
    this.getLocation();
  }

  ngOnInit(): void {
  }

  getLocation(): void {
    this.sHttpService.getLocation(this.data.idLocation).subscribe(
      (location) => {
        this.location = location;
        this.location.residents.forEach(ch => {
          let id = ch;
          if (ch.length > 2) {
            const url = ch.split('/');
            id = url[url.length - 1];
          }
          this.getNameCharacter(id);
        });
        this.status = 'ok';
      },
      (err) => {

      }
    );
  }

  getNameCharacter(id: string): void {
    this.sHttpService.getCharacter(id).subscribe(
      (character) => {
        this.characters.push(character.name);
      },
      (err) => {
      }
    );
  }


}
