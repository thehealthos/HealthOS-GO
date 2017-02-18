/*
 * healthos_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ) on 02/18/2017
 */

package diseases_pkg


/*
 * Interface for the DISEASES_IMPL
 */
type DISEASES interface {
    GetDisease (string) (interface{}, error)

    GetAllDiseases (int64, int64) (interface{}, error)

    SearchDiseases (string) (interface{}, error)
}

/*
 * Factory for the DISEASES interaface returning DISEASES_IMPL
 */
func NewDISEASES() DISEASES {
    return &DISEASES_IMPL{}
}