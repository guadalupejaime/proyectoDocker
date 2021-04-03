import { Component, OnInit } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { PageEvent } from '@angular/material/paginator';
import { MatSnackBar } from '@angular/material/snack-bar';
import { MatTableDataSource } from '@angular/material/table';
import { Character } from '@pages/dashboard/characters/characters.interface';
import { AddCharacterComponent } from '@pages/dashboard/characters/modals/add-character/add-character.component';
import { InfoCharacterComponent } from '@pages/dashboard/characters/modals/info-character/info-character.component';
import { TablePage } from '@services/http.interfaces';
import { HttpService } from '@services/http.service';

@Component({
  selector: 'app-characters',
  templateUrl: './characters.component.html',
  styleUrls: ['./characters.component.scss']
})
export class CharactersComponent implements OnInit {
  status = 'load';
  displayedColumns: string[] = ['name', 'status', 'specie', 'gender', 'info'];
  totalCharacters = 0;
  characters: Character[] = [];
  dataCharacters = new MatTableDataSource<Character>(this.characters);
  sizePage = 20;
  paginatorSettings: TablePage;

  constructor(
    private sHttpService: HttpService,
    private sMatDialog: MatDialog,
    private sMatSnackBar: MatSnackBar
  ) { }

  ngOnInit(): void {
    this.paginatorSettings = {
      limit: this.sizePage,
      offset: 0
    };

    this.sHttpService.getCharacters(this.paginatorSettings).subscribe(
      (characters) => {
        this.totalCharacters = characters.total_found;
        this.dataCharacters.data = characters.results;
        // this.dataCharacters = characters.results;
        this.status = 'ok';
      },
      (err) => {
        this.status = 'error';
      }
    );
  }

  /** update characters table */
  updateCharacters(): void {
    this.sHttpService.getCharacters(this.paginatorSettings).subscribe(
      (characters) => {
        this.totalCharacters = characters.total_found;
        this.dataCharacters.data = characters.results;
      },
      (err) => {
      }
    );
  }

  /** update paginato settings */
  updatePaginatorSettings(event: PageEvent): void {
    this.paginatorSettings = {
      limit: event.pageSize,
      offset: event.pageIndex * event.pageSize
    };
    this.updateCharacters();
  }

  /** search on data table */
  applyFilter(event: Event): void {
    const filterValue = (event.target as HTMLInputElement).value;
    this.dataCharacters.filter = filterValue.trim().toLowerCase();
  }

  /** open a modal wt character information selected */
  infoCharacter(id: number): void {
    const modal = this.sMatDialog.open(InfoCharacterComponent, {
      maxHeight: '99%',
      maxWidth: '99%',
      disableClose: true,
      panelClass: 'myModal',
      backdropClass: 'myBackdrop',
      data: { idCharacter: id }
    });
  }


  /** open a  modal to add a new character */
  addCharacter(): void {
    const modal = this.sMatDialog.open(AddCharacterComponent, {
      maxHeight: '99%',
      maxWidth: '99%',
      disableClose: true,
      panelClass: 'myModal',
      backdropClass: 'myBackdrop',
    });

    modal.afterClosed().subscribe(
      (result) => {
        if (result) {
          this.openNotification('Added new character, success!');
          this.updateCharacters();
          return;
        } else if (result === 'error') {
          this.openNotification('Something went wrong!');
        }
      });
  }

  openNotification(message: string): void {
    this.sMatSnackBar.open(message, 'End now', {
      duration: 1500,
      horizontalPosition: 'right',
      verticalPosition: 'top',
    });
  }

}
