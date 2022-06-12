package rs.edu.raf.banka.racun.dto;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@NoArgsConstructor
@AllArgsConstructor
public class AkcijePodaciDto {
    String ticker;
    String opisHartije;
    String time;
    Double open;
    Double high;
    Double low;
    Double price;
    Long volume;
    Double previousClose;
    Double change;
    Double changePercent;
    Long outstandingShares;
    Long berzaId;
    Long id;

    // Calculated
    Double priceVolume;
    Double marketCap;

    public Double getPriceVolume() {
        return volume * price;
    }

    public Double getMarketCap() {
        return outstandingShares * price;
    }

}
