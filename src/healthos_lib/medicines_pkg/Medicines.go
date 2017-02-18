/*
 * healthos_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ) on 02/18/2017
 */

package medicines_pkg


/*
 * Interface for the MEDICINES_IMPL
 */
type MEDICINES interface {
    GetManufacturer (string) (interface{}, error)

    SearchManufacturers (string) (interface{}, error)

    GetCommonSideEffects (string) (interface{}, error)

    GetPopularUsage (string) (interface{}, error)

    GetMedicinesByManufacturer (int64, int64, string) (interface{}, error)

    GetSimilarMedicines (string) (interface{}, error)

    GetMedicine (string) (interface{}, error)

    GetAllMedicines (int64, int64) (interface{}, error)

    SearchMedicines (string) (interface{}, error)

    GetAllManufacturers (int64, int64) (interface{}, error)

    GetAlternativeMedicines (int64, int64, string) (interface{}, error)
}

/*
 * Factory for the MEDICINES interaface returning MEDICINES_IMPL
 */
func NewMEDICINES() MEDICINES {
    return &MEDICINES_IMPL{}
}