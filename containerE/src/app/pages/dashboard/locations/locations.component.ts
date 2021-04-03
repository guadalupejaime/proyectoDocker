import { Component, OnInit } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { PageEvent } from '@angular/material/paginator';
import { MatSnackBar } from '@angular/material/snack-bar';
import { MatTableDataSource } from '@angular/material/table';
import { Location } from '@pages/dashboard/locations/locations.interface';
import { AddLocationComponent } from '@pages/dashboard/locations/modals/add-location/add-location.component';
import { InfoLocationComponent } from '@pages/dashboard/locations/modals/info-location/info-location.component';
import { TablePage } from '@services/http.interfaces';
import { HttpService } from '@services/http.service';

@Component({
  selector: 'app-locations',
  templateUrl: './locations.component.html',
  styleUrls: ['./locations.component.scss']
})
export class LocationsComponent implements OnInit {
  status = 'load';
  displayedColumns: string[] = ['name', 'type', 'dimension', 'info'];
  totalLocations = 0;
  locations: Location[] = [];
  dataLocations = new MatTableDataSource<Location>(this.locations);
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

    this.sHttpService.getLocations(this.paginatorSettings).subscribe(
      (locations) => {
        this.totalLocations = locations.total_found;
        this.dataLocations.data = locations.results;
        // this.dataLocations = locations.results;
        this.status = 'ok';
      },
      (err) => {
        this.status = 'error';
      }
    );
  }

  /** update locations table */
  updateLocations(): void {
    this.sHttpService.getLocations(this.paginatorSettings).subscribe(
      (locations) => {
        this.totalLocations = locations.total_found;
        this.dataLocations.data = locations.results;
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
    this.updateLocations();
  }

  /** search on data table */
  applyFilter(event: Event): void {
    const filterValue = (event.target as HTMLInputElement).value;
    this.dataLocations.filter = filterValue.trim().toLowerCase();
  }

  /** open a modal wt character information selected */
  infoLocation(id: number): void {
    const modal = this.sMatDialog.open(InfoLocationComponent, {
      maxHeight: '99%',
      maxWidth: '99%',
      disableClose: true,
      panelClass: 'myModal',
      backdropClass: 'myBackdrop',
      data: { idLocation: id }
    });
  }


  /** open a  modal to add a new character */
  addLocation(): void {
    const modal = this.sMatDialog.open(AddLocationComponent, {
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
