/*
 * healthos_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ) on 02/18/2017
 */

package generics_pkg


/*
 * Interface for the GENERICS_IMPL
 */
type GENERICS interface {
    GetGeneric (string) (interface{}, error)

    GetMedicinesByGeneric (int64, int64, bool, string) (interface{}, error)

    GetAllGenerics (int64, int64) (interface{}, error)

    SearchGenerics (string) (interface{}, error)
}

/*
 * Factory for the GENERICS interaface returning GENERICS_IMPL
 */
func NewGENERICS() GENERICS {
    return &GENERICS_IMPL{}
}