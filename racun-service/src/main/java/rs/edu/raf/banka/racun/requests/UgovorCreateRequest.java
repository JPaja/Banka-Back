package rs.edu.raf.banka.racun.requests;

import lombok.Data;

@Data
public class UgovorCreateRequest
{
    Long companyId;

    String delodavniBroj;

    String description;

}
